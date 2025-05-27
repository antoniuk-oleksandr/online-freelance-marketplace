interface Google {
    accounts: {
        id: {
            initialize: (config: { client_id: string, callback: (response: any) => void }) => void,
            prompt: () => void,
        },
        oauth2: {
            initCodeClient: (config: {
                client_id: string,
                scope: string,
                ux_mode: string,
                callback: (response: any) => void,
                error_callback: (response: any) => void,
            }) => { requestCode: () => void },
            initTokenClient: (config: {
                client_id: string,
                scope: string,
                ux_mode: string,
                callback: (response: any) => void,
                error_callback: (response: any) => void,
            }) => { requestAccessToken: () => void },
        },
    },
}

declare const google: Google;

interface Window {
    env: {
        VITE_GOOGLE_CLIENT_ID: string,
        VITE_MAX_FREELANCE_BY_ID_REVIEWS: string,
        VITE_SERVICE_FEES: string,
        VITE_FILE_SERVER_HOST: string,
        VITE_WEBSOCKET_HOST: string,
        VITE_BACKEND_HOST: string,
    }
}

declare const window: Window;


