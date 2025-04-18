import { request } from "@/api/request"
import { errorStore } from "@/common-stores/error-store"
import { themeStore } from "@/common-stores/theme-storage"
import type { DropdownItem } from "@/types/DropdownItem"
import type { PostSignOutRequestResponse } from "@/types/PostSignOutRequestResponse"
import type { SignHeaderData } from "@/types/SignHeaderData"
import Cookies from "js-cookie"
import { navigate } from "svelte-routing"
import { signDataStore } from "./components/HeaderProfileBlock/sign-data-store"
import { toastElementStore } from "../ToastElement/store/toast-element-store"

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


export const getUserSession = (): SignHeaderData => {
  const accessToken = Cookies.get('accessToken')
  const refreshToken = Cookies.get('refreshToken')
  const avatar = localStorage.getItem('accessTokenAvatar')
  const userId = localStorage.getItem('accessTokenUserId')

  let signedIn = true
  if (!refreshToken || avatar === undefined || userId === undefined) {
    Cookies.remove('refreshToken')
    Cookies.remove('accessToken')
    localStorage.removeItem('accessTokenAvatar')
    localStorage.removeItem('accessTokenUserId')

    signedIn = false
  }

  return { avatar, signedIn, userId, accessToken, refreshToken }
}

const logoutClickAction = async (signData: SignHeaderData) => {
  const logoutBody = { accessToken: signData.accessToken, refreshToken: signData.refreshToken }
  const { data, status } = await request<PostSignOutRequestResponse>("POST", "/auth/sign-out", logoutBody)

  if (status !== 200) {
    errorStore.set({ shown: true, error: data.error })
    return
  }

  signDataStore.update(() => ({ ...signData, signedIn: false }))
  toastElementStore.update((prev) => ({ ...prev, show: true, message: "You have successfully signed out.", type: "success" }))

  Cookies.remove('refreshToken')
  Cookies.remove('accessToken')

  localStorage.removeItem("accessTokenAvatar");
  localStorage.removeItem("accessTokenUsername");
  localStorage.removeItem("accessTokenUserId");
  navigate('/')
}
