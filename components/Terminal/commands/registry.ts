export type State = { filesystem: { cwd: string } };
export type Params = string[];
export type CommandHandler = (
  state: State,
  params: Params
) => string | undefined;
type Command = { fn: CommandHandler; help?: string };
export const registry: Record<string, Command> = {};

type registerParams = { name: string; fn: CommandHandler; help?: string };

export default function register({ name, fn, help }: registerParams) {
  registry[name] = { fn, help };
}

export function get_command(name: string) {
  return registry[name]?.fn;
}

export function get_help(name: string) {
  return registry[name]?.help
}

export function get_all_commands() {
  return Object.keys(registry)
}
