import { baseUrl } from '@/constants/constants';
import type { components } from '@/types/openapi/screening';
import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: baseUrl,
  withCredentials: true,
});

export const getScreenings = async (day?: string) => {
  const response = await axiosInstance.get('/screenings/', {
    params: day ? { day } : undefined,
  });
  return response.data as components['schemas']['ScreeningOutput'][];
};

export const postCreateScreening = async (
  data: components['schemas']['CreateScreeningInput']
) => {
  const response = await axiosInstance.post('/screenings/', data);
  return response.data as components['schemas']['ScreeningOutput'];
};

export const getScreeningById = async (id: string) => {
  const response = await axiosInstance.get(`/screenings/${id}`);
  return response.data as components['schemas']['ScreeningOutput'];
};

export const putUpdateScreeningById = async (
  id: string,
  data: { startTime: string }
) => {
  await axiosInstance.put(`/screenings/${id}`, data);
};

export const deleteScreeningById = async (id: string) => {
  await axiosInstance.delete(`/screenings/${id}`);
};

export const getScreeningsByMovie = async (movieID: string, day?: string) => {
  const response = await axiosInstance.get(`/screenings/movie/${movieID}`, {
    params: day ? { day } : undefined,
  });
  return response.data as components['schemas']['ScreeningOutput'][];
};

export const getScreeningsByRoom = async (roomID: string, day?: string) => {
  const response = await axiosInstance.get(`/screenings/room/${roomID}`, {
    params: day ? { day } : undefined,
  });
  return response.data as components['schemas']['ScreeningOutput'][];
};

export const getScreeningsByCinema = async (cinemaID: string, day?: string) => {
  const response = await axiosInstance.get(`/screenings/cinema/${cinemaID}`, {
    params: day ? { day } : undefined,
  });
  return response.data as components['schemas']['ScreeningOutput'][];
};
