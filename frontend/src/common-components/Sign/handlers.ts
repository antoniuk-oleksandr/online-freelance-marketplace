import {ResponseErrorEnum} from "@/types/ResponseErrorEnum.ts";
import {ResponseMessageEnum} from "@/types/ResponseMessageEnum.ts";
import Cookies from "js-cookie";
import {navigate} from "svelte-routing";
import {toastElementStore} from "@/common-components/ToastElement/store/toast-element-store.ts";
import {postGoogleRequest} from "@/api/post-google-request.ts";

const showToast = (message: string, type: "success" | "error") => {
    toastElementStore.set({
        show: true,
        message,
        type,
        exitAnimation: false,
    });
};

const setCookies = (data: any, keepSignedIn: boolean, values: any) => {
    let options: ({ expires: number } | undefined)[] = [undefined, undefined];
    if (keepSignedIn) {
        options[0] = {expires: 15 / (24 * 60)}; // 15 minutes
        options[1] = {expires: 30}; // 30 days
    }

    Cookies.set("accessToken", data.accessToken, options[0]);
    Cookies.set("refreshToken", data.refreshToken, options[1]);
};

const handleSignErrors = (data: any, setErrors: any, setFields: any) => {
    if (data.error === ResponseErrorEnum.UsernameIsTaken) {
        setErrors("username", "Username is already taken.");
    } else if (data.error === ResponseErrorEnum.EmailIsTaken) {
        setErrors("email", "Email is already taken.");
    } else if (data.error === ResponseErrorEnum.InvalidEmail) {
        setErrors("email", "Invalid email address.");
    } else if (data.error === ResponseErrorEnum.InvalidCredentials) {
        setFields("password", "");
        setTimeout(() => {
            setErrors("usernameOrEmail", "Invalid email/username or password.");
            setErrors("password", "Invalid email/username or password.");
        });
    }
};

const handleSuccessSign = (data: any, values: any, setShowEmailSentMessage: any) => {
    if (data.message === ResponseMessageEnum.EmailSentSuccessfully) {
        setShowEmailSentMessage && setShowEmailSentMessage(true);
    } else {
        setCookies(data, values.keepSignedIn, values);
        navigate("/");
        showToast("You have successfully signed in.", "success");
    }
};

export const handleSignSubmit = async (
    values: any,
    submitAction: (values: any) => Promise<any>,
    setErrors: (fieldName: string, error: string) => void,
    setFields: (fieldName: string, error: string) => void,
    setShowEmailSentMessage?: (showEmailSentMessage: boolean) => void,
    setLoading?: (loading: boolean) => void,
) => {
    if (!setLoading) return;
    setLoading(true);

    const {data, status} = await submitAction(values);

    if (status === 200) {
        handleSuccessSign(data, values, setShowEmailSentMessage);
    } else {
        handleSignErrors(data, setErrors, setFields);
    }

    setLoading(false);
};


export const handleGoogleAuth = async (code: string) => {
    const {status, data} = await postGoogleRequest(code);
    if (status === 200) {
        handleSuccessSign(data, {keepSignedIn: false}, false);
    }
}

export const handleGoogleButtonClick = async (
    setLoading: (loading: boolean) => void,
    clientId: string,
) => {
    setLoading(true);
    google.accounts.oauth2.initCodeClient({
        client_id: clientId,
        scope: "email profile",
        ux_mode: "popup",
        callback: async (googleResponse) => {
            await handleGoogleAuth(googleResponse.code);
            setLoading(false);
        },
        error_callback: () => setLoading(false)
    }).requestCode();
};