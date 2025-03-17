import { Interaction, SlashCommandBuilder } from "discord.js";

export default {
  data: new SlashCommandBuilder()
    .setName("ping")
    .setDescription("Checks if bot is running."),

  async execute(interaction: Interaction) {
    if (!interaction.isChatInputCommand()) {
      return;
    }

    await interaction.reply("pong!");
  },
};
