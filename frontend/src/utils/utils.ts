export const getHost = () => {
    return 'localhost:8080';
}

export const getFileHost = () => {
    return 'localhost:8030';
}

export const getFile = (name: string) => {
    if (name.includes("http")) return name;
    else return `http://${getFileHost()}/files/${name}`;
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
            return `${Math.floor(diffInDays / 365)} years ago`;
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
