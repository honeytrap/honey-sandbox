import { createClientAction } from '../shared/helpers/createClientAction';

export const Connect = createClientAction('connection.connect');
export const Connected = createClientAction<{ connected: boolean, errors: string }>('connection.connected');
export const ErrorAction = createClientAction<{ message: string, sendToServer: true }>('connection.error');
export const CloseError = createClientAction('connection.closeError');
export const LoadingDone = createClientAction('connection.loadingDone');
