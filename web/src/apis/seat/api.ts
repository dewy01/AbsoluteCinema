import { baseUrl } from '@/constants/constants';
import type { components } from '@/types/openapi/seat';
import axios from 'axios';

export const axiosInstance = axios.create({
  baseURL: baseUrl,
  withCredentials: true,
});

export const postCreateSeat = async (
  data: components['schemas']['CreateSeatInput']
) => {
  const response = await axiosInstance.post('/seats/', data);
  return response.data as components['schemas']['SeatOutput'];
};

export const getSeatById = async (id: string) => {
  const response = await axiosInstance.get(`/seats/${id}`);
  return response.data as components['schemas']['SeatOutput'];
};

export const putUpdateSeatById = async (
  id: string,
  data: components['schemas']['UpdateSeatInput']
) => {
  await axiosInstance.put(`/seats/${id}`, data);
};

export const deleteSeatById = async (id: string) => {
  await axiosInstance.delete(`/seats/${id}`);
};

export const getSeatsByRoomId = async (roomID: string) => {
  const response = await axiosInstance.get(`/seats/room/${roomID}`);
  return response.data as components['schemas']['SeatOutput'][];
};

export const getSeatsByScreeningId = async (screeningID: string) => {
  const response = await axiosInstance.get(`/seats/screening/${screeningID}`);
  return response.data as components['schemas']['SeatWithReservationStatusOutput'][];
};
