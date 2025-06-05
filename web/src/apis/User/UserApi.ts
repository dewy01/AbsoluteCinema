import { baseUrl } from '@/constants/constants';
import type { components } from '@/types/openapi/user';
import axios from 'axios';

export interface AccessToken {
  accessToken: string;
  refreshToken: string;
}

export const axiosInstance = axios.create({ baseURL: baseUrl, withCredentials: true });

export const postUserRegister = async (data: components['schemas']['CreateUserInput']) => {
  return await axiosInstance.post('/users/register', data);
};

export const postUserLogin = async (data: components['schemas']['LoginUserInput']) => {
  return await axiosInstance.post('/users/login', data);
};

export const postUserLogout = async () => {
  return await axiosInstance.post('/users/logout');
};

export const getCurrentUser = async () => {
  const response = await axiosInstance.get('/users/me');
  return response.data as components['schemas']['UserOutput'];
};

export const getUserById = async (id: string) => {
  const response = await axiosInstance.get(`/users/${id}`);
  return response.data as components['schemas']['UserOutput'];
};

export const putUserById = async (id: string, data: components['schemas']['UpdateUserInput']) => {
  const response = await axiosInstance.put(`/users/${id}`, data);
  return response.data as components['schemas']['UserOutput'];
};

export const deleteUserById = async (id: string) => {
  return await axiosInstance.delete(`/users/${id}`);
};
