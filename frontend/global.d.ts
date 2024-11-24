interface Google {
    accounts: {
        id: {
            initialize: (config: { client_id: string; callback: (response: any) => void }) => void;
            prompt: () => void;
        };
        oauth2: {
            initCodeClient: (config: {
                client_id: string;
                scope: string;
                ux_mode: string;
                callback: (response: any) => void;
                error_callback: (response: any) => void;
            }) => { requestCode: () => void };
        };
    };
}

declare const google: Google;