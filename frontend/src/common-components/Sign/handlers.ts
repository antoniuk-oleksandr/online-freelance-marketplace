import { postAuthRequest } from "@/api/post-auth-request";
import { toastElementStore } from "@/common-components/ToastElement/store/toast-element-store";
import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";
import { ResponseMessageEnum } from "@/types/ResponseMessageEnum";
import { navigate } from "svelte-routing";
import { setTokenCookies } from "./helpers";

const showToast = (message: string, type: "success" | "error") => {
    toastElementStore.set({
        show: true,
        message,
        type,
    });
};


const handleSignErrors = (data: any, setErrors: any, setFields: any) => {
    switch (data.error) {
        case ResponseErrorEnum.UsernameIsTaken:
            setErrors("username", "Username is already taken.");
            break;
        case ResponseErrorEnum.EmailIsTaken:
            setErrors("email", "Email is already taken.");
            break;
        case ResponseErrorEnum.InvalidEmail:
            setErrors("email", "Invalid email address.");
            break;
        case ResponseErrorEnum.InvalidCredentials:
            setFields("password", "");
            setTimeout(() => {
                setErrors("usernameOrEmail", "Invalid email/username or password.");
                setErrors("password", "Invalid email/username or password.");
            });
            break;
        case ResponseErrorEnum.EmailDoesNotExist:
            setErrors("usernameOrEmail", "Email does not exist.");
            break;
        case ResponseErrorEnum.UsernameDoesNotExist:
            setErrors("usernameOrEmail", "Username does not exist.");
            break;
    }
};


const handleSuccessSign = (
    backendResponse: any,
    formData: any,
    setShowEmailSentMessage: any,
) => {
    if (backendResponse.message === ResponseMessageEnum.EmailSentSuccessfully) {
        setShowEmailSentMessage && setShowEmailSentMessage(true);
    } else if (backendResponse.accessToken && backendResponse.refreshToken) {
        setTokenCookies(backendResponse, formData.keepSignedIn);
        navigate("/");
        showToast("You have successfully signed in.", "success");
    }
};

export const handleSignSubmit = async (
    formData: any,
    submitAction: (values: any) => Promise<any>,
    setErrors: (fieldName: string, error: string) => void,
    setFields: (fieldName: string, error: string) => void,
    setShowEmailSentMessage?: (showEmailSentMessage: boolean) => void,
    setLoading?: (loading: boolean) => void,
) => {
    if (!setLoading) return;
    setLoading(true);

    const { data: responseBackend, status } = await submitAction(formData);

    if (status === 200) {
        handleSuccessSign(responseBackend, formData, setShowEmailSentMessage);
    } else {
        handleSignErrors(responseBackend, setErrors, setFields);
    }

    setLoading(false);
};


export const handleGoogleAuth = async (code: string, keepSignedIn?: boolean) => {
    const { status, data: backendResponse } = await postAuthRequest("google", undefined, { code });

    if (status === 200) {
        handleSuccessSign(backendResponse, { keepSignedIn }, false);
    }
}

export const handleGoogleButtonClick = async (
    setLoading: (loading: boolean) => void,
    clientId: string,
    keepSignedIn?: boolean
) => {
    setLoading(true);
    google.accounts.oauth2.initCodeClient({
        client_id: clientId,
        scope: "email profile",
        ux_mode: "popup",
        callback: async (googleResponse) => {
            await handleGoogleAuth(googleResponse.code, keepSignedIn);
            setLoading(false);
        },
        error_callback: () => setLoading(false)
    }).requestCode();
};
