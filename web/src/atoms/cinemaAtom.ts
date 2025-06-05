import { atom } from 'jotai';

const localStorageKey = 'selectedCinema';

const getInitialCinema = () => {
  try {
    const stored = localStorage.getItem(localStorageKey);
    return stored ? JSON.parse(stored) : null;
  } catch {
    return null;
  }
};

export const selectedCinemaAtom = atom(getInitialCinema());

selectedCinemaAtom.onMount = (setAtom) => {
  const handleStorage = (event: StorageEvent) => {
    if (event.key === localStorageKey) {
      setAtom(event.newValue ? JSON.parse(event.newValue) : null);
    }
  };
  window.addEventListener('storage', handleStorage);
  return () => window.removeEventListener('storage', handleStorage);
};

export const writeSelectedCinemaAtom = atom(
  null,
  (get, set, newCinema: null | { id?: string; name?: string; address?: string; roomIDs?: string[] }) => {
    if (newCinema) {
      localStorage.setItem(localStorageKey, JSON.stringify(newCinema));
    } else {
      localStorage.removeItem(localStorageKey);
    }
    set(selectedCinemaAtom, newCinema);
  }
);
