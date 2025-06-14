import { baseUrl } from '@/constants/constants';
import type { components } from '@/types/openapi/cinema';
import axios from 'axios';

export const axiosInstance = axios.create({ baseURL: baseUrl, withCredentials: true });

export const getCinemas = async () => {
  const response = await axiosInstance.get('/cinemas/');
  return response.data as components['schemas']['CinemaOutput'][];
};

export const postCreateCinema = async (data: components['schemas']['CreateCinemaInput']) => {
  const response = await axiosInstance.post('/cinemas/', data);
  return response.data as components['schemas']['CinemaOutput'];
};

export const getCinemaById = async (id: string) => {
  const response = await axiosInstance.get(`/cinemas/${id}`);
  return response.data as components['schemas']['CinemaOutput'];
};

export const putUpdateCinemaById = async (
  id: string,
  data: components['schemas']['UpdateCinemaInput']
) => {
  const response = await axiosInstance.put(`/cinemas/${id}`, data);
  return response.data as components['schemas']['CinemaOutput'];
};

export const deleteCinemaById = async (id: string) => {
  return await axiosInstance.delete(`/cinemas/${id}`);
};
