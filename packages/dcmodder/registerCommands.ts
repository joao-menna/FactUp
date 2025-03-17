import { REST, Routes } from "discord.js";

const commands = [
  {
    name: "ping",
    description: "Checks if bot is running.",
  },
];

const token = process.env.DCMODDER_BOT_TOKEN;
const clientId = process.env.DCMODDER_CLIENT_ID;

async function main() {
  if (!token || !clientId) {
    throw new Error("I need a token/clientId to work!");
  }

  const rest = new REST().setToken(token);

  try {
    console.log("Started refreshing application slash commands.");

    await rest.put(Routes.applicationCommands(clientId), { body: commands });

    console.log("Successfully reloaded application slash commands.");
  } catch (err) {
    console.error(err);
  }
}

main();
