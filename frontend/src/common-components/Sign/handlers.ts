import { request } from "@/api/request";
import { toastElementStore } from "@/common-components/ToastElement/store/toast-element-store";
import { errorStore } from "@/common-stores/error-store";
import { setPrivateKey } from "@/common-stores/private-key-store";
import type { GoogleUserInfo } from "@/types/GoogleUserInfo";
import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";
import { ResponseMessageEnum } from "@/types/ResponseMessageEnum";
import type { SignInData } from "@/types/SignInData";
import type { SignInRequestResponse } from "@/types/SignInRequestResponse";
import { decryptWithKey, encryptWithKey, generateAESIV, generateAESSalt } from "@/utils/aes-utils";
import { setIndexedDbItem } from "@/utils/indexed-db-utils";
import { navigate } from "svelte-routing";
import { signDataStore } from "../Header/components/HeaderProfileBlock/sign-data-store";
import { encryptUserCredentials, makeRequestToCheckIfUserExists, makeRequestToGetGoogleUserInfo, parseSignInUserBackendResponse } from "./helpers";
import { processChatPartners, processSharedSecrets } from "../Header/helpers";
import { setSharedSecrets } from "@/common-stores/shared-secrets-store";

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


const handleSuccessSign = async (
    response: SignInData,
    formData: any,
    setShowEmailSentMessage: any,
) => {
    //@ts-ignore
    if (response.message === ResponseMessageEnum.EmailSentSuccessfully) {
        setShowEmailSentMessage && setShowEmailSentMessage(true);
        return;
    }

    localStorage.setItem("accessTokenAvatar", response.userData.avatar);
    localStorage.setItem("accessTokenUserId", response.userData.id.toString());

    const encryptionData = parseSignInUserBackendResponse(response.userData);

    const decryptedPrivateKey = await decryptWithKey(
        encryptionData.privateKey, formData.password,
        encryptionData.privateKeyIV, encryptionData.privateKeySalt
    );

    setPrivateKey(decryptedPrivateKey)

    const privateKeyIV = generateAESIV();
    const privateKeySalt = generateAESSalt();
    const encryptedPrivateKey = await encryptWithKey(
        decryptedPrivateKey, encryptionData.masterKey, privateKeyIV, privateKeySalt
    );

    await setIndexedDbItem('privateKey', encryptedPrivateKey);
    await setIndexedDbItem('privateKeyIV', privateKeyIV);
    await setIndexedDbItem('privateKeySalt', privateKeySalt);

    signDataStore.set({
        avatar: response.userData.avatar,
        userId: response.userData.id.toString(),
        authenticated: true,
    })


    const chatPartners = processChatPartners(response.chatPartners);
    const sharedSecrets = await processSharedSecrets(chatPartners, decryptedPrivateKey);
    setSharedSecrets(sharedSecrets)

    navigate("/");
    showToast("You have successfully signed in.", "success");
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
        await handleSuccessSign(responseBackend, formData, setShowEmailSentMessage);
    } else {
        handleSignErrors(responseBackend, setErrors, setFields);
    }

    setLoading(false);
};


export const handleGoogleAuth = async (
    userInfo: GoogleUserInfo, keepSignedIn?: boolean, accessToken?: string
) => {
    const { available, error } = await makeRequestToCheckIfUserExists(userInfo.email)
    if (error) return;

    let response = {} as SignInRequestResponse
    if (available) {
        //sign up
        const body = await encryptUserCredentials({ keepSignedIn, accessToken }, userInfo.sub);
        response = await request<SignInRequestResponse>("PUT", "/auth/google", body, true)

    } else {
        //sign in
        response = await request<SignInRequestResponse>(
            "POST", "/auth/google", { keepSignedIn, accessToken }, true
        );
    }
    if (response.status !== 200) errorStore.set({ shown: true, error: response.data.error })
    else await handleSuccessSign(response.data, { keepSignedIn, password: userInfo.sub }, false);
}

//here it starts
export const handleGoogleButtonClick = async (
    setLoading: (loading: boolean) => void,
    clientId: string,
    keepSignedIn?: boolean
) => {
    setLoading(true);
    google.accounts.oauth2.initTokenClient({
        client_id: clientId,
        scope: 'openid',
        ux_mode: 'redirect',
        callback: async (response) => {
            try {
                const userInfo = await makeRequestToGetGoogleUserInfo(response.access_token);
                await handleGoogleAuth(userInfo, keepSignedIn, response.access_token);
            } catch (error) {
                console.error('Error processing Google auth:', error);
            } finally {
                setLoading(false);
            }
        },
        error_callback: (error) => {
            console.error('Google OAuth error:', error);
            setLoading(false);
        }
    }).requestAccessToken();
};

