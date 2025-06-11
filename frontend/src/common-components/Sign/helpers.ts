import { request } from "@/api/request";
import { errorStore } from "@/common-stores/error-store";
import type { GetEmailAvailabilityRequestResponse } from "@/types/GetEmailAvailabilityRequestResponse";
import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";
import type { SignInData } from "@/types/SignInData";
import type { SignInEncryptionData } from "@/types/SignInEncryptionData";
import type { UserData } from "@/types/UserData";
import { encryptWithKey, generateAESIV } from "@/utils/aes-utils";
import { convertBase64ToUint8Array } from "@/utils/base64-utils";
import { generateECDHKeyPair } from "@/utils/ecdh-utils";
import axios from "axios";

export const parseSignInUserBackendResponse = (
  userData: UserData,
): SignInEncryptionData => {
  return {
    privateKey: convertBase64ToUint8Array(userData.privateKey),
    privateKeyIV: convertBase64ToUint8Array(userData.privateKeyIV),
    privateKeySalt: convertBase64ToUint8Array(userData.privateKeySalt),
    masterKey: convertBase64ToUint8Array(userData.masterKey)
  }
}

export const makeRequestToCheckIfUserExists = async (
  email: string,
): Promise<{ available: boolean; error: boolean; }> => {
  const url = `/auth/email-availability?email=${email}`
  const response = await request<GetEmailAvailabilityRequestResponse>("GET", url);
  if (response.status === 200) {
    return {
      ...response.data,
      error: false
    };
  } else {
    errorStore.set({ shown: true, error: response.data.error })
    return {
      available: false,
      error: true
    }
  }
}

export const makeRequestToGetGoogleUserInfo = async (accessToken: string) => {
  const url = 'https://www.googleapis.com/oauth2/v3/userinfo'

  try {
    const response = await axios.get(url, {
      headers: { Authorization: `Bearer ${accessToken}` }
    });

    if (response.status === 200) return response.data;
    else errorStore.set({ shown: true, error: ResponseErrorEnum.UnexpectedError });
    return null;
  }
  catch (error) {
    errorStore.set({ shown: true, error: ResponseErrorEnum.UnexpectedError });
    return null;
  }
}

export const encryptUserCredentials = async (body: any, key: string | Uint8Array) => {
  const keys = await generateECDHKeyPair()

  const privateKeyIV = generateAESIV();
  const privateKeySalt = generateAESIV();
  const encryptedKey = await encryptWithKey(keys.privateKey, key, privateKeyIV, privateKeySalt);

  return {
    ...body,
    privateKeyIV: Array.from(privateKeyIV),
    privateKeySalt: Array.from(privateKeySalt),
    privateKey: Array.from(encryptedKey),
    publicKey: Array.from(keys.publicKey),
  }
}
