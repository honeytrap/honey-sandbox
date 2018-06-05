import { action, base } from 'ts-action';

export function createClientAction<T extends {} = {}>(type: string) {
	return action(type, base(
		class {
			public sendToServer = false;
			constructor(public payload?: T) {}
		}
	));
}