import {writable} from "svelte/store";

const initialTheme = localStorage.getItem('theme') === 'dark';

export const isDarkMode = writable(initialTheme);

isDarkMode.subscribe(value => {
    if (value) {
        document.documentElement.classList.add('dark');
        localStorage.setItem('theme', 'dark');
    } else {
        document.documentElement.classList.remove('dark');
        localStorage.setItem('theme', 'light');
    }
})