import type { Category } from "@/types/Category";
import { SvelteURLSearchParams } from "svelte/reactivity";

export const getHost = () => {
    return window.env.VITE_BACKEND_HOST    
    // return 'http://localhost:8080';
}

export const getWSHost = () => {
    return window.env.VITE_WEBSOCKET_HOST
}

export const getFileServerHost = () => {
    return window.env.VITE_FILE_SERVER_HOST
    // return 'http://localhost:8030/files';
}

export const getHoverClass = () => {
    return "hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover duration-200 ease-out"
}

export const getTimeAgo = (date: number): string => {
    const now = new Date();
    const diffInMs = now.getTime() - new Date(date).getTime();
    const diffInDays = Math.floor(diffInMs / (1000 * 60 * 60 * 24));

    switch (true) {
        case diffInDays === 0:
            return "today";
        case diffInDays === 1:
            return "yesterday";
        case diffInDays < 7:
            return `${diffInDays} days ago`;
        case diffInDays < 14:
            return "1 week ago";
        case diffInDays < 21:
            return "2 weeks ago";
        case diffInDays < 30:
            return "3 weeks ago";
        case diffInDays < 60:
            return "1 month ago";
        case diffInDays < 365:
            return `${Math.floor(diffInDays / 30)} months ago`;
        default:
            const rest = Math.floor(diffInDays / 365);
            return `${rest} year${rest === 1 ? '' : 's'} ago`;
    }
};

export const getPackageDurationDifference = (before: number, after: number) => {
    const beforeData = new Date(before);
    const afterData = new Date(after);
    const diffInMs = afterData.getTime() - beforeData.getTime();
    const diffInDays = Math.floor(diffInMs / (1000 * 60 * 60 * 24));

    switch (true) {
        case diffInDays < 1:
            return `${Math.floor(diffInMs / (1000 * 60 * 60))} hours`;
        case diffInDays === 1:
            return "1 day";
        case diffInDays < 7:
            return `${diffInDays} days`;
        case diffInDays === 7:
            return "1 week";
        case diffInDays < 30:
            return `${Math.floor(diffInDays / 7)} weeks`;
        case diffInDays === 30:
            return "1 month";
        case diffInDays < 365:
            return `${Math.floor(diffInDays / 30)} months`;
        case diffInDays === 365:
            return "1 year";
        default:
            return `${Math.floor(diffInDays / 365)} years`;
    }
}

export const capitalize = (str: string) => {
    return str.charAt(0).toUpperCase() + str.slice(1);
}

export const formatText = (text: string) => {
    return text.replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

export const round = (value: number, precision: number) => {
    const multiplier = Math.pow(10, precision || 0);
    return Math.round((value + Number.EPSILON) * multiplier) / multiplier;
};


export const getServiceFees = () => {
    return parseFloat(window.env.VITE_SERVICE_FEES as string);
}

export const calcPriceWithServiceFees = (value: number) => {
    const serviceFees = getServiceFees();
    return round(value + value * serviceFees, 2);
}

export const flyFade = (node: Element, { x = 0, y = 0, duration = 300 }) => {
    return {
        duration,
        css: (t: number) => `
        opacity: ${t};
        transform: translate(${(1 - t) * x}px, ${(1 - t) * y}px);
      `,
    }
}

export const getPageParam = () => {
    const params = new SvelteURLSearchParams(window.location.search);
    return parseInt(params.get('page') || '1');
}

export const getUserId = (): number | null => {
    const idStr = localStorage.getItem("accessTokenUserId");
    return idStr !== null ? parseInt(idStr) : null
}

export const getServiceCategories = (): Category[] => [
    { id: 1, name: 'Web Development' },
    { id: 2, name: 'Graphic Design' },
    { id: 3, name: 'Digital Marketing' },
    { id: 4, name: 'SEO Services' },
    { id: 5, name: 'Content Writing' },
    { id: 6, name: 'Mobile App Development' },
    { id: 7, name: 'Data Analysis' },
    { id: 8, name: 'Cybersecurity' },
    { id: 9, name: 'Cloud Computing' },
    { id: 10, name: 'E-commerce Development' },
    { id: 11, name: 'UI/UX Design' },
    { id: 12, name: 'Video Production' },
    { id: 13, name: 'Photography' },
    { id: 14, name: 'Virtual Assistance' },
    { id: 15, name: 'Software Development' },
    { id: 16, name: 'Business Consulting' },
    { id: 17, name: 'IT Support' },
    { id: 18, name: 'Project Management' },
    { id: 19, name: 'Copywriting' },
    { id: 20, name: 'Social Media Management' }
];

export const getServiceCategoryIdByName = (name: string): number => {
    const categories = getServiceCategories();
    const category = categories.find((category) => category.name === name);
    return category ? category.id : -1;
}

export const convertFileArrToFileList = (fileArr: File[]): FileList => {
    const dataTransfer = new DataTransfer();
    fileArr.forEach((file) => {
        dataTransfer.items.add(file);
    })

    return dataTransfer.files;
}
