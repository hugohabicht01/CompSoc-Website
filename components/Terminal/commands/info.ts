import register, { CommandHandler, registry } from "./registry";
import { whoami } from "./session";

const echo: CommandHandler = (_, params) => {
  return params.join(" ");
};

const info: CommandHandler = (_, params) => {
  let target = params.join(" ");
  switch (target.toLowerCase()) {
    case "compsoc":
    case "lucompsoc":
    case "computer science society":
      return "Lancaster University Computer Science Society exists to promote interest in computing and technology among students and wider society.";

    case "lu":
    case "uni":
    case "lancaster university":
      return "Lancaster University is a collegiate university 3 miles south of Lancaster.";

    default:
      return `Could not find information for \`${target}\`. Try \`info compsoc\``;
  }
};

const neofetch: CommandHandler = (state, _) => {
  const responseLines = [];
  responseLines.push(`${whoami(state, [])}@compsoc.io`);
  responseLines.push(
    "-".repeat((responseLines.at(responseLines.length - 1) as string).length)
  );
  responseLines.push("OS: CompSocOS (Terminal Edition)");
  responseLines.push("Shell: holy-sea 0.0.1");
  responseLines.push("Resolution: 80x25");
  responseLines.push("Theme: Matrix-red");
  responseLines.push("Terminal: megantereon");
  responseLines.push("Processor: unknown");
  return responseLines.join("\n");
};

const help: CommandHandler = (state, _) => {
  const commandStrings = Object.keys(registry).join("\n");
  return `Help:\nThe following commands are available:\n${commandStrings}`;
};

register("echo", echo);
register("info", info);
register("neofetch", neofetch);
register("help", help);
