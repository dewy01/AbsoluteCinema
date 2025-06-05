import { useAllCinemas } from '@/apis/Cinema';
import type { CinemaMap } from '@/types/CinemaMap';
import { createContext, useContext, useState } from 'react';

interface CinemaContextProps {
  selectedCinema: string | null;
  isCinemaModalOpen: boolean;
  cinemaMap: CinemaMap;
  selectCinema: (cinema: string) => void;
  resetCinema: () => void;
  openCinemaModal: () => void;
  closeCinemaModal: () => void;
}

const cinemaContext = createContext<CinemaContextProps | undefined>(undefined);

type Props = { children: React.ReactNode };

export const CinemaProvider = ({ children }: Props) => {
  const [selectedCinema, setSelectedCinema] = useState(() => {
    return localStorage.getItem('selectedCinema');
  });

  const [isCinemaModalOpen, setCinemaModalOpen] = useState(false);

  const selectCinema = (cinema: string) => {
    localStorage.setItem('selectedCinema', cinema);
    setSelectedCinema(cinema);
    setCinemaModalOpen(false);
  };

  const resetCinema = () => {
    localStorage.removeItem('selectedCinema');
    setSelectedCinema(null);
  };

  const openCinemaModal = () => setCinemaModalOpen(true);
  const closeCinemaModal = () => setCinemaModalOpen(false);

  const { data } = useAllCinemas();
  const cinemaMap: CinemaMap = (data || []).reduce((acc, cinema) => {
    if (cinema.id) {
      acc[cinema.id] = {
        name: cinema.name || '',
        address: cinema.address || '',
        roomIDs: cinema.roomIDs || []
      };
    }
    return acc;
  }, {} as CinemaMap);

  return (
    <cinemaContext.Provider
      value={{
        selectedCinema,
        cinemaMap,
        isCinemaModalOpen,
        selectCinema,
        resetCinema,
        openCinemaModal,
        closeCinemaModal
      }}>
      {children}
    </cinemaContext.Provider>
  );
};

export const useCinema = () => {
  const context = useContext(cinemaContext);
  if (!context) {
    throw new Error('useCinema must be used within a CinemaProvider');
  }
  return context;
};
