import {
  Client,
  Collection,
  Events,
  GatewayIntentBits,
  Interaction,
} from "discord.js";
import { readdirSync } from "fs";
import { join } from "path";

const commands = new Collection();

async function loadCommands() {
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
        commands.set(command.data.name, command);
      } else {
        console.log(
          `[WARNING] The command at ${filePath} is missing a required "data" or "execute" property.`
        );
      }
    }
  }
}

async function main() {
  await loadCommands();

  const client = new Client({
    intents: [GatewayIntentBits.Guilds],
  });

  client.on(Events.ClientReady, (readyClient) => {
    console.log(`Logged in as ${readyClient.user.tag}`);
  });

  client.on(Events.InteractionCreate, async (interaction) => {
    if (!interaction.isChatInputCommand()) {
      return;
    }

    const command = commands.get(interaction.commandName) as {
      execute: (interaction: Interaction) => Promise<void>;
    };

    if (!command) {
      console.error(
        `No command matching ${interaction.commandName} was found.`
      );
      return;
    }

    try {
      await command.execute(interaction);
    } catch (err) {
      console.error(err);
    }
  });

  client.login(process.env.DCMODDER_BOT_TOKEN);
}

main();
