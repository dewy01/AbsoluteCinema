import { createContext, useContext, useState } from 'react';

interface CinemaContextProps {
  selectedCinema: string | null;
  isCinemaModalOpen: boolean;
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

  return (
    <cinemaContext.Provider
      value={{
        selectedCinema,
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
