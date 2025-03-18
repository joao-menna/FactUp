import { REST, Routes } from "discord.js";
import { readdirSync } from "fs";
import { join } from "path";

const token = process.env.DCMODDER_BOT_TOKEN;
const clientId = process.env.DCMODDER_CLIENT_ID;

async function getCommands() {
  const commands = [];

  const foldersPath = join(__dirname, "commands");
  const commandFolders = readdirSync(foldersPath);

  for (const folder of commandFolders) {
    const commandsPath = join(foldersPath, folder);
    const commandFiles = readdirSync(commandsPath).filter((file) =>
      file.endsWith(".js")
    );
    for (const file of commandFiles) {
      const filePath = join(commandsPath, file);
      const command = (await import(filePath)).default;
      if ("data" in command && "execute" in command) {
        commands.push(command.data.toJSON());
      } else {
        console.log(
          `[WARNING] The command at ${filePath} is missing a required "data" or "execute" property.`
        );
      }
    }
  }

  return commands;
}

async function main() {
  if (!token || !clientId) {
    throw new Error("I need a token/clientId to work!");
  }

  const commands = await getCommands();

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
