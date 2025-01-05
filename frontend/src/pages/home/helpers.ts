const defineSlice = (width: number) => {
    switch (true) {
        case width < 640:
            return 2;
        case width < 768:
            return 3;
        case width < 1024:
            return 4;
        case width < 1280:
            return 5;
        case width < 1536:
            return 6;
        default:
            return 6;
    }
}

export const sliceHomeCategories = (categories: any) => {
    return categories.slice(0, defineSlice(window.innerWidth));
}

export const homeCategories = [
    {
        title: "Web Development",
        icon: "hugeicons:code",
        id: 1,
    },
    {
        title: "Mobile App Development",
        icon: "hugeicons:smart-phone-01",
        id: 6,
    },
    {
        title: "Graphic Design",
        icon: "hugeicons:brush",
        id: 2,
    },
    {
        title: "Digital Marketing",
        icon: "hugeicons:megaphone-02",
        id: 3,
    },
    {
        title: "Video Production",
        icon: "hugeicons:film-roll-02",
        id: 12,
    },
    {
        title: "Photography",
        icon: "hugeicons:camera-02",
        id: 13,
    }
];


