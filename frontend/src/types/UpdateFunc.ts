export type Updater<T> = (prev: T) => Promise<T>;
export type UpdateFunc<T> = (updater: Updater<T>) => void;