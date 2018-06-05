import { action, base } from 'ts-action';

export function createServerAction<T>(type: string) {
	return action(type, base(
		class {
			public sendToServer = true;
			constructor(public payload: T) {}
		}
	));
}