export const getHost = () => {
    return 'http://localhost:8080';
}

export const getFileServerHost = () => {
    return 'http://localhost:8030';
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

export const getPackageDuration = (before: number, after: number) => {
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
    return parseFloat(import.meta.env.VITE_SERVICE_FEES as string);
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
