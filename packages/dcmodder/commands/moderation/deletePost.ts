import { Interaction, SlashCommandBuilder } from "discord.js";

export default {
  data: new SlashCommandBuilder()
    .setName("deletepost")
    .setDescription("Deletes a post using PostID.")
    .addIntegerOption((b) =>
      b
        .setName("postid")
        .setDescription("PostID to delete")
        .setMinValue(1)
        .setRequired(true)
    ),

  async execute(interaction: Interaction) {
    if (!interaction.isChatInputCommand()) {
      return;
    }

    const baseUrl = process.env.DCMODDER_BACKEND_URL ?? "http://factup_proxy";
    const postId = interaction.options.getInteger("postid");
    const url = `${baseUrl}/api/v1/post/${postId}`;
    console.log(url);

    try {
      const req = await fetch(url, {
        method: "DELETE",
        headers: { authorization: process.env.DCMODDER_BACKEND_TOKEN ?? "" },
      });

      console.log(await req.text());

      if (req.status !== 200) {
        throw new Error("could not delete post");
      }
    } catch (err) {
      await interaction.reply(`Couldn't delete post ${postId}!`);
      return;
    }

    await interaction.reply(`Post ${postId} has been deleted!`);
  },
};
