import type { components } from "@/types/openapi/screening";

export const mapScreeningsByMovie = (screenings: components['schemas']['ScreeningOutput'][]) => {
  const grouped: Record<
    string,
    {
      movie: components['schemas']['MovieOutput'];
      screenings: {
        id?: string;
        startTime?: string;
        room?: components['schemas']['RoomOutput'];
      }[];
    }
  > = {};

  screenings.forEach((screening) => {
    const movieId = screening.movie?.id;
    if (!movieId) return;

    if (!grouped[movieId]) {
      grouped[movieId] = {
        movie: screening.movie!,
        screenings: []
      };
    }

    grouped[movieId].screenings.push({
      id: screening.id,
      startTime: screening.startTime,
      room: screening.room
    });
  });

  return Object.values(grouped);
};