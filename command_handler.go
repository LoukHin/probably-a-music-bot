package main

import (
	"context"
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgolink/v2/lavalink"
	"github.com/disgoorg/log"
)

func (b *Bot) shuffle(event *events.ApplicationCommandInteractionCreate, _ discord.SlashCommandInteractionData) error {
	queue := b.Guilds.GetQueue(*event.GuildID())
	if queue == nil {
		return updateInteractionResponse(event, "No player found")
	}

	queue.Shuffle()
	return updateInteractionResponse(event, "Queue shuffled")
}

func (b *Bot) volume(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return updateInteractionResponse(event, "No player found")
	}

	volume := data.Int("volume")
	if err := player.Update(context.TODO(), lavalink.WithVolume(volume)); err != nil {
		return updateInteractionResponse(event, fmt.Sprintf("Error while setting volume: `%s`", err))
	}

	return updateInteractionResponse(event, fmt.Sprintf("Volume set to `%d`", volume))
}

func (b *Bot) seek(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return updateInteractionResponse(event, "No player found")
	}

	position := data.Int("position")
	unit, ok := data.OptInt("unit")
	if !ok {
		unit = 1
	}
	finalPosition := lavalink.Duration(position * unit)
	if err := player.Update(context.TODO(), lavalink.WithPosition(finalPosition)); err != nil {
		return updateInteractionResponse(event, fmt.Sprintf("Error while seeking: `%s`", err))
	}

	return updateInteractionResponse(event, fmt.Sprintf("Seeked to `%s`", formatDuration(finalPosition)))
}

func (b *Bot) skip(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	queue := b.Guilds.GetQueue(*event.GuildID())
	if player == nil || queue == nil {
		return updateInteractionResponse(event, "No player found")
	}

	amount, ok := data.OptInt("amount")
	if !ok {
		amount = 1
	}

	track, ok := queue.Skip(amount)
	if !ok {
		return updateInteractionResponse(event, "No tracks in queue")
	}

	if err := player.Update(context.TODO(), lavalink.WithTrack(track)); err != nil {
		return updateInteractionResponse(event, fmt.Sprintf("Error while skipping track: `%s`", err))
	}

	return updateInteractionResponse(event, "Skipped track")
}

func (b *Bot) repeatType(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	queue := b.Guilds.GetQueue(*event.GuildID())
	if queue == nil {
		return updateInteractionResponse(event, "No player found")
	}

	queue.Type = QueueType(data.String("mode"))
	queue.RecalculateDuration()
	b.updatePlayerMessage(*event.GuildID())
	return updateInteractionResponse(event, fmt.Sprintf("Repeat mode set to `%s`", queue.Type))
}

func (b *Bot) clearQueue(event *events.ApplicationCommandInteractionCreate, _ discord.SlashCommandInteractionData) error {
	queue := b.Guilds.GetQueue(*event.GuildID())
	if queue == nil {
		return updateInteractionResponse(event, "No player found")
	}

	queue.Clear()
	return updateInteractionResponse(event, "Queue cleared")
}

func (b *Bot) removeQueue(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	queue := b.Guilds.GetQueue(*event.GuildID())
	if queue == nil {
		return updateInteractionResponse(event, "No player found")
	}

	removedTrack, ok := queue.Remove(data.Int("id") - 1)
	if !ok {
		return updateInteractionResponse(event, "Can't remove track")
	}
	return updateInteractionResponse(event, fmt.Sprintf("Removed [%s](%s)", removedTrack.Info.Title, *removedTrack.Info.URI))
}

func (b *Bot) queue(event *events.ApplicationCommandInteractionCreate, _ discord.SlashCommandInteractionData) error {
	queue := b.Guilds.GetQueue(*event.GuildID())
	if queue == nil {
		return updateInteractionResponse(event, "No player found")
	}

	if len(queue.Tracks) == 0 {
		return updateInteractionResponse(event, "No tracks in queue")
	}

	var tracks string
	for i, track := range queue.Tracks {
		tracks += fmt.Sprintf("%d. [`%s`](<%s>)\n", i+1, track.Info.Title, *track.Info.URI)
	}

	return updateInteractionResponse(event, fmt.Sprintf("Queue `%s`:\n%s", queue.Type, tracks))
}

func (b *Bot) pause(event *events.ApplicationCommandInteractionCreate, _ discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return updateInteractionResponse(event, "No player found")
	}

	if err := player.Update(context.TODO(), lavalink.WithPaused(!player.Paused())); err != nil {
		return updateInteractionResponse(event, fmt.Sprintf("Error while pausing: `%s`", err))
	}

	status := "playing"
	if player.Paused() {
		status = "paused"
	}
	return updateInteractionResponse(event, fmt.Sprintf("Player is now %s", status))
}

func (b *Bot) stop(event *events.ApplicationCommandInteractionCreate, _ discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return updateInteractionResponse(event, "No player found")
	}

	if err := player.Update(context.TODO(), lavalink.WithNullTrack()); err != nil {
		return updateInteractionResponse(event, fmt.Sprintf("Error while stopping: `%s`", err))
	}

	return updateInteractionResponse(event, "Player stopped")
}

func (b *Bot) disconnect(event *events.ApplicationCommandInteractionCreate, _ discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return updateInteractionResponse(event, "No player found")
	}

	if ok := b.updateVoiceState(*event.GuildID(), nil); !ok {
		return updateInteractionResponse(event, "Error while disconnecting")
	}

	return updateInteractionResponse(event, "Player disconnected")
}

func (b *Bot) nowPlaying(event *events.ApplicationCommandInteractionCreate, _ discord.SlashCommandInteractionData) error {
	player := b.Lavalink.ExistingPlayer(*event.GuildID())
	if player == nil {
		return updateInteractionResponse(event, "No player found")
	}

	track := player.Track()
	if track == nil {
		return updateInteractionResponse(event, "No track found")
	}

	return updateInteractionResponse(event, fmt.Sprintf("Now playing: [`%s`](<%s>)\n\n %s / %s", track.Info.Title, *track.Info.URI, formatDuration(player.Position()), formatDuration(track.Info.Length)))
}

func (b *Bot) play(event *events.ApplicationCommandInteractionCreate, data discord.SlashCommandInteractionData) error {
	query := data.String("query")

	var err error
	b.playOrQueue(*event.GuildID(), event.Member().Member, query, func(embed discord.Embed) {
		_, err = event.Client().Rest().UpdateInteractionResponse(event.ApplicationID(), event.Token(), discord.NewMessageUpdateBuilder().SetEmbeds(embed).Build())
	})
	return err
}

func (b *Bot) setup(event *events.ApplicationCommandInteractionCreate, _ discord.SlashCommandInteractionData) error {
	var (
		message *discord.Message
		embed   discord.EmbedBuilder
	)

	guild, err := getEntClient().Guild.Get(context.TODO(), *event.GuildID())
	if err != nil {
		return err
	}

	if guild.PlayerChannelID == nil || *guild.PlayerChannelID != event.ChannelID() {
		embed.SetTitle("Nothing currently playing")
		embed.SetImage("https://images.pexels.com/videos/3045163/free-video-3045163.jpg?auto=compress&cs=tinysrgb&dpr=1")
		message, err = b.Client.Rest().CreateMessage(event.ChannelID(), discord.NewMessageCreateBuilder().SetEmbeds(embed.Build()).SetContent("Join a voice channel and queue songs by name or url in here.").Build())
		if err != nil {
			log.Error(err)
		}
		guild, err = guild.Update().SetPlayerChannelID(event.ChannelID()).SetPlayerMessageID(message.ID).Save(context.TODO())
		if err != nil {
			log.Error(err)
		}
		return updateInteractionResponse(event, "Player created")
	} else {
		return updateInteractionResponse(event, "This channel was already a player!")
	}
}