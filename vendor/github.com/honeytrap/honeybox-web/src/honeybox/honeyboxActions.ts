import { createServerAction } from '../shared/helpers/createServerAction';
import { createClientAction } from '../shared/helpers/createClientAction';

export const StdIn = createServerAction<{ data: string }>('stdin');
export const StdOut = createClientAction<{ data: string }>('stdout');
export const WriteStdOut = createClientAction<{ data: string }>('honeybox.writeStdOut');
export const Event = createClientAction<{ data: string }>('event');
export const StdErr = createClientAction<{ data: string }>('stderr');
export const SetMessageFilter = createClientAction<{ query: string }>('honeybox.setMessageFilter');
export const Info = createClientAction<{ name: string }>('info');
