import { baseUrl } from '@/constants/constants';
import type { components } from '@/types/openapi/room';
import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: baseUrl,
  withCredentials: true,
});

export const postCreateRoom = async (
  data: components['schemas']['CreateRoomInput']
) => {
  const response = await axiosInstance.post('/rooms/', data);
  return response.data as components['schemas']['RoomOutput'];
};

export const getRoomById = async (id: string) => {
  const response = await axiosInstance.get(`/rooms/${id}`);
  return response.data as components['schemas']['RoomOutput'];
};

export const putUpdateRoomById = async (
  id: string,
  data: components['schemas']['UpdateRoomInput']
) => {
  await axiosInstance.put(`/rooms/${id}`, data);
};

export const deleteRoomById = async (id: string) => {
  await axiosInstance.delete(`/rooms/${id}`);
};

export const getRoomsByCinema = async (cinemaID: string) => {
  const response = await axiosInstance.get(`/rooms/cinema/${cinemaID}`);
  return response.data as components['schemas']['RoomOutput'][];
};
