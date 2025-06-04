import { baseUrl } from '@/constants/constants';
import { MutationCache, QueryCache, QueryClient } from '@tanstack/react-query';
import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: baseUrl,
  withCredentials: true
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
  if (window.location.pathname === '/connection') {
    window.location.href = '/';
  }
};

const handleErrorCache = async (err: unknown) => {
  if (axios.isAxiosError(err)) {
    const status = err.response?.status;

    if (status === 401) {
      // queryClient.removeQueries({ queryKey: ['me'] }); to powoduje petle

      // const publicPaths = ['/login', '/register'];
      // if (!publicPaths.includes(window.location.pathname)) {
      //   window.location.href = '/login';
      // }

      return;
    }

    if (status === 403) {
      window.location.href = '/forbidden';
      return;
    }

    if (status === 404) {
      window.location.href = '/not-found';
      return;
    }
  }

  console.error('Unhandled error:', err);
};
