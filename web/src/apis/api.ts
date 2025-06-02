import { baseUrl } from '@/constants/constants';
import { MutationCache, QueryCache, QueryClient } from '@tanstack/react-query';
import { type AccessToken, axiosInstance as noAuthInstance } from './User/UserApi';
import axios from 'axios';

export const axiosInstance = axios.create({ baseURL: baseUrl });

axiosInstance.interceptors.request.use((config) => {
  const authToken = localStorage.getItem('authToken');
  if (authToken) {
    config.headers['Authorization'] = `Bearer ${authToken}`;
  }
  return config;
});

export const queryClient = new QueryClient({
  mutationCache: new MutationCache({
    onSuccess: () => {
      handleSuccessCache();
    },
    onError: async (err) => {
      await handleErrorCache(err);
    }
  }),
  queryCache: new QueryCache({
    onSuccess: () => {
      handleSuccessCache();
    },
    onError: async (err) => {
      await handleErrorCache(err);
    }
  }),
  defaultOptions: {
    mutations: {
      retry: false
    },
    queries: {
      retry: false
    }
  }
});

const handleSuccessCache = () => {
  if (location.pathname === '/connection') {
    location.pathname = '/';
  }
};
const handleErrorCache = async (err: unknown) => {
  if (axios.isAxiosError(err)) {
    if (err.response?.status === 401) {
      try {
        await refreshAccessToken();

        // try to re-execute
        if (err.config) {
          if (err.config.method) {
            return await axiosInstance(err.config);
          }
        }

        return;
      } catch (error) {
        localStorage.removeItem('accessToken');
        localStorage.removeItem('refreshToken');
        queryClient.removeQueries();
        location.reload();
      }
    } else if (err.response?.status === 403) {
      location.pathname = '/forbidden';
    } else if (err.response?.status === 404) {
      location.pathname = '/not-found';
    }
  } else {
    console.error('An unexpected error occurred:', err);
  }
};

export const refreshAccessToken = async () => {
  const refToken = localStorage.getItem('refreshToken');
  const authToken = localStorage.getItem('accessToken');

  try {
    const response = await noAuthInstance.post('/api/token/refresh', {
      accessToken: authToken,
      refreshToken: refToken
    } as AccessToken);
    const { accessToken, refreshToken } = response.data as AccessToken;

    localStorage.setItem('accessToken', accessToken);
    localStorage.setItem('refreshToken', refreshToken);
    return;
  } catch (error) {
    throw new Error('Failed to refresh access token');
  }
};
