import type { RefreshTokenRequestResponse } from "@/types/RefreshTokenRequestResponse";
import { ResponseErrorEnum } from "@/types/ResponseErrorEnum";
import axios from "axios";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

export const getAccessToken = async (host: string): Promise<RefreshTokenRequestResponse> => {
  const accessToken = Cookies.get("accessToken");
  if (accessToken) {
    const decodedJWT = jwtDecode(accessToken);
    const expired = new Date(decodedJWT.exp as number * 1000) < new Date()
    if (!expired) return { data: { accessToken }, status: 200 };
  }

  const response = await tryToRefreshToken(host);
  if (response.status === 200) {
    const decodedData = jwtDecode(response.data.accessToken);
    Cookies.set("accessToken", response.data.accessToken, { expires: decodedData.exp });
    return response;
  } else {
    return response;
  }
}

export const tryToRefreshToken = async (host: string): Promise<RefreshTokenRequestResponse> => {
  const refreshToken = Cookies.get("refreshToken");
  if (!refreshToken) return {
    data: { error: ResponseErrorEnum.ExpiredToken },
    status: 401,
  }

  const url = `${host}/api/v1/auth/refresh-token`;

  try {
    const headers = { Authorization: `Bearer ${refreshToken}` };
    return await axios.post(url, undefined, { headers });
  } catch (e) {
    
    return {
      data: (e as any).response.data,
      status: (e as any).status,
    }
  }
}
