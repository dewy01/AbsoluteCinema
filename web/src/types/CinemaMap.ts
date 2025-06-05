export interface CinemaMap {
  [cinemaId: string]: {
    name?: string;
    address?: string;
    roomIDs?: string[];
  };
}
