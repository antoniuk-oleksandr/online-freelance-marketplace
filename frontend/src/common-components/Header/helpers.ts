import { request } from "@/api/request"
import { errorStore } from "@/common-stores/error-store"
import { setPrivateKey } from "@/common-stores/private-key-store"
import { themeStore } from "@/common-stores/theme-storage"
import type { DropdownItem } from "@/types/DropdownItem"
import type { GetUserSessionRequestResponse } from "@/types/GetUserSessionRequestResponse"
import type { PostSignOutRequestResponse } from "@/types/PostSignOutRequestResponse"
import type { SignHeaderData } from "@/types/SignHeaderData"
import type { UserSessionData } from "@/types/UserSessionData"
import { decryptWithKey } from "@/utils/aes-utils"
import { convertBase64ToUint8Array } from "@/utils/base64-utils"
import { clearIndexedDbItem, getIndexedDbItem } from "@/utils/indexed-db-utils"
import { navigate } from "svelte-routing"
import { toastElementStore } from "../ToastElement/store/toast-element-store"
import { signDataStore } from "./components/HeaderProfileBlock/sign-data-store"
import { ChatMessageType } from "@/types/ChatMessageType"
import type { ChatPartnerPublicKeyData } from "@/types/ChatPartnerPublicKeyData"
import type { ChatPartnerSecret } from "@/types/ChatPartnerSecret"
import { deriveSharedSecret } from "@/utils/ecdh-utils"
import { setSharedSecrets } from "@/common-stores/shared-secrets-store"

export const getMyProfileDropdownItems = (
  signData: SignHeaderData, darkMode: boolean | null,
): DropdownItem[] => [
    {
      title: 'View Profile',
      clickAction: () => navigate(`/users/${signData.userId}`),
      icon: 'hugeicons:user'
    },
    {
      title: 'Edit Profile',
      clickAction: () => navigate('/my-profile/edit'),
      icon: 'hugeicons:user-edit-01',
      dividerAfter: true
    },
    {
      title: 'My Orders',
      clickAction: () => navigate('/my-profile/orders'),
      icon: 'hugeicons:package'
    },
    {
      title: 'My Services',
      clickAction: () => navigate('/my-profile/services'),
      icon: 'hugeicons:briefcase-03',
    },
    {
      title: 'Incoming Requests',
      clickAction: () => navigate('/my-profile/requests'),
      icon: 'hugeicons:task-01',
      dividerAfter: true
    },
    {
      title: 'Theme',
      clickAction: () => { themeStore.set(darkMode === null ? false : !darkMode) },
      badge: darkMode ? 'Dark' : 'Light',
      icon: `hugeicons:${darkMode ? 'moon-02' : 'sun-03'}`,
      dividerAfter: true,
      closeDropdown: false,
    },
    {
      title: 'Log Out',
      clickAction: () => logoutClickAction(signData),
      icon: 'hugeicons:door-02',
      customColor: 'text-red-500 dark:text-red-500',
    },
  ]


export const getUserSession = async (): Promise<SignHeaderData> => {
  try {
    const avatar = localStorage.getItem('accessTokenAvatar')
    const userId = localStorage.getItem('accessTokenUserId')

    if (avatar === null || userId === null) {
      return { avatar: "", userId: "", authenticated: false }
    }

    let userSessionData = await processSessionRequest();
    if (!userSessionData.authenticated) {
      return { avatar: "", userId: "", authenticated: false }
    }

    userSessionData = processUserSessionData(userSessionData);

    const encryptedPrivateKey = await getIndexedDbItem("privateKey");
    const privateKeyIV = await getIndexedDbItem("privateKeyIV");
    const privateKeySalt = await getIndexedDbItem("privateKeySalt");
    const decryptedPrivateKey = await decryptWithKey(
      encryptedPrivateKey, userSessionData.masterKey, privateKeyIV, privateKeySalt
    );
    setPrivateKey(decryptedPrivateKey);

    const sharedSecrets = await processSharedSecrets(userSessionData.chatPartners, decryptedPrivateKey);
    setSharedSecrets(sharedSecrets)

    return {
      avatar, userId,
      authenticated: userSessionData.authenticated,
      chatPartners: userSessionData.chatPartners
    };
  }
  catch (error) {
    return {
      avatar: "",
      userId: "",
      authenticated: false,
    }
  }
}

export const processSharedSecrets = async (
  chatPartners: ChatPartnerPublicKeyData[],
  privateKey: Uint8Array
): Promise<ChatPartnerSecret[]> => Promise.all(
  chatPartners.map(async (item) => ({
    userId: item.userId,
    secret: await deriveSharedSecret(privateKey, item.publicKey as Uint8Array)
  }))
)

const processUserSessionData = (userSessionData: UserSessionData) => ({
  ...userSessionData,
  masterKey: convertBase64ToUint8Array(userSessionData.masterKey as string),
  chatPartners: processChatPartners(userSessionData.chatPartners)
})

export const processChatPartners = (chatPartners: ChatPartnerPublicKeyData[]) =>
  chatPartners.map((partner) => ({
    ...partner,
    publicKey: convertBase64ToUint8Array(partner.publicKey as string),
  }))


const logoutClickAction = async (signData: SignHeaderData) => {
  const { data, status } = await request<PostSignOutRequestResponse>(
    "POST", "/auth/sign-out", undefined, true
  );

  if (status !== 200) {
    errorStore.set({ shown: true, error: data.error });
    return;
  }

  signDataStore.update(() => ({
    ...signData, authenticated: false, masterKey: "", masterKeyIv: ""
  }));
  toastElementStore.update((prev) => ({
    ...prev, show: true, message: "You have successfully signed out.", type: "success"
  }));

  localStorage.removeItem("accessTokenAvatar");
  localStorage.removeItem("accessTokenUsername");
  localStorage.removeItem("accessTokenUserId");

  await clearIndexedDbItem("privateKey");
  await clearIndexedDbItem("privateKeyIV");
  await clearIndexedDbItem("privateKeySalt");
  setPrivateKey(new Uint8Array(0));
  navigate('/');
}

const processSessionRequest = async (): Promise<UserSessionData> => {
  try {
    const response = await request<GetUserSessionRequestResponse>(
      "GET", "/auth/session", undefined, true
    );
    if (response.status === 200) return {
      ...response.data,
      authenticated: true,
    };
  }
  catch (error) { }

  return {
    authenticated: false,
    masterKey: "",
    chatPartners: []
  }
}
