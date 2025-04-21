import {writable} from "svelte/store";

const initialTheme = localStorage.getItem('theme') === 'dark';

export const themeStore = writable(initialTheme);

themeStore.subscribe(value => {
    if (value) {
        document.documentElement.classList.add('dark');
        localStorage.setItem('theme', 'dark');
        document.documentElement.setAttribute('data-theme', 'dark');
    } else {
        document.documentElement.classList.remove('dark');
        localStorage.setItem('theme', 'light');
        document.documentElement.setAttribute('data-theme', 'light');
    }
})
