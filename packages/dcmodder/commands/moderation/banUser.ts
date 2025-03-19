import { Interaction, SlashCommandBuilder } from "discord.js";

export default {
  data: new SlashCommandBuilder()
    .setName("banuser")
    .setDescription("Bans an user from using the platform.")
    .addIntegerOption((b) =>
      b
        .setName("userid")
        .setDescription("UserID to ban")
        .setMinValue(1)
        .setRequired(true)
    ),

  async execute(interaction: Interaction) {
    if (!interaction.isChatInputCommand()) {
      return;
    }

    const baseUrl = process.env.DCMODDER_BACKEND_URL ?? "http://factup_proxy";
    const userId = interaction.options.getInteger("userid");
    const url = `${baseUrl}/api/v1/user/ban/${userId}`;
    console.log(url);

    try {
      const req = await fetch(url, {
        method: "DELETE",
        headers: { authorization: process.env.DCMODDER_BACKEND_TOKEN ?? "" },
      });

      console.log(await req.text());

      if (req.status !== 200) {
        throw new Error("could not ban user");
      }
    } catch (err) {
      await interaction.reply(`Couldn't ban user ${userId}!`);
      return;
    }

    await interaction.reply(`User ${userId} has been banned!`);
  },
};
