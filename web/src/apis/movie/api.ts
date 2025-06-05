import { baseUrl } from '@/constants/constants';
import type { components } from '@/types/openapi/movie';
import axios from 'axios';

export const axiosInstance = axios.create({ baseURL: baseUrl, withCredentials: true });

export const getMovies = async () => {
  const response = await axiosInstance.get('/movies/');
  return response.data as components['schemas']['MovieOutput'][];
};

export const postCreateMovie = async (data: components['schemas']['CreateMovieInput']) => {
  const formData = new FormData();
  formData.append('title', data.title);
  formData.append('director', data.director);
  if (data.description) formData.append('description', data.description);
  if (data.actorIDs) data.actorIDs.forEach((id) => formData.append('actorIDs', id));
  formData.append('photo', data.photo);

  const response = await axiosInstance.post('/movies/', formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  });
  return response.data as components['schemas']['MovieOutput'];
};

export const getMovieById = async (id: string) => {
  const response = await axiosInstance.get(`/movies/${id}`);
  return response.data as components['schemas']['MovieOutput'];
};

export const putUpdateMovieById = async (id: string, data: components['schemas']['UpdateMovieInput']) => {
  const formData = new FormData();
  if (data.title) formData.append('title', data.title);
  if (data.director) formData.append('director', data.director);
  if (data.description) formData.append('description', data.description);
  if (data.actorIDs) data.actorIDs.forEach((id) => formData.append('actorIDs', id));
  if (data.photo) formData.append('photo', data.photo);

  const response = await axiosInstance.put(`/movies/${id}`, formData, {
    headers: { 'Content-Type': 'multipart/form-data' },
  });
  return response.data as components['schemas']['MovieOutput'];
};

export const deleteMovieById = async (id: string) => {
  return await axiosInstance.delete(`/movies/${id}`);
};
