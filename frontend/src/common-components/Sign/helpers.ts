import type { AccessToken } from "@/types/AccessToken";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

const decodeAccessToken = (token: string) => {
  const decoded = jwtDecode<AccessToken>(token);
  
  localStorage.setItem("accessTokenAvatar", decoded.avatar);
  localStorage.setItem("accessTokenUsername", decoded.username);
  localStorage.setItem("accessTokenUserId", decoded.userId);
  return decoded;
};

const getTokenExpiration = (token: string): Date | undefined => {
  const decoded = jwtDecode<{ exp: number }>(token);
  return decoded.exp ? new Date(decoded.exp * 1000) : undefined;
};

export const setTokenCookies = (data: { accessToken: string; refreshToken: string }, keepSignedIn: boolean) => {
  decodeAccessToken(data.accessToken);

  let options: ({ expires: Date } | undefined)[] = [undefined, undefined];


  if (keepSignedIn) {
    const accessTokenExpiration = getTokenExpiration(data.accessToken);
    const refreshTokenExpiration = getTokenExpiration(data.refreshToken);

    if (accessTokenExpiration && refreshTokenExpiration) {
      options = [{ expires: accessTokenExpiration }, { expires: refreshTokenExpiration }];
    }
  }

  Cookies.set("accessToken", data.accessToken, options[0]);
  Cookies.set("refreshToken", data.refreshToken, options[1]);
};
