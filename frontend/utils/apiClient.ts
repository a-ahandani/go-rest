import axios from "axios";
import createAuthRefreshInterceptor from "axios-auth-refresh";

export const apiClient = axios.create({
    baseURL: 'http://localhost:3000/api',
    withCredentials: true,
})
export const setHeaderToken = (token: string) => {
    apiClient.defaults.headers.common.Authorization = `Bearer ${token}`;
};

export const removeHeaderToken = () => {
    delete apiClient.defaults.headers.common.Authorization;
};

const fetchNewToken = async () => {
    try {
        const token: string = await apiClient
            .get("/refresh-token")
            .then(res => res.data.token);
        return token;
    } catch (error) {
        return null;
    }
};

const refreshAuth = async (failedRequest: any) => {
    const newToken = await fetchNewToken();

    if (newToken) {
        failedRequest.response.config.headers.Authorization = "Bearer " + newToken;
        setHeaderToken(newToken);
        return Promise.resolve(newToken);
    } else {
        // router.push("/login");
        return Promise.reject();
    }
};


createAuthRefreshInterceptor(apiClient, refreshAuth, {
    statusCodes: [401], // default: [ 401 ]
    pauseInstanceWhileRefreshing: true,
});

