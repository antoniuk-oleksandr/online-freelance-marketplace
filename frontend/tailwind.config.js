import typography from '@tailwindcss/typography'


/** @type {import('tailwindcss').Config} */
const config = {
    content: [
        "./index.html",
        "./src/**/*.svelte",
    ],
    safelist: [
        "lg:min-w-header-dropdown-menu",
        "size-12", "size-64", "size-14", "size-24",
        "size-32", "lg:min-w-48", "lg:min-w-60",
        "size-48",
        "lg:size-12", "lg:size-64", "lg:size-14", "size-16",
        "bg-gray-200", "text-gray-700", "border-gray-300",
        "bg-yellow-100", "text-yellow-700", "border-yellow-300",
        "bg-green-100", "text-green-700", "border-green-300",
        "bg-red-200", "text-red-800", "border-red-400",
        "bg-orange-100", "text-orange-700", "border-orange-300",
        "bg-red-100", "text-red-700", "border-red-300",
        "bg-blue-100", "text-blue-700", "border-blue-300"
    ],
    theme: {
        extend: {
            transitionProperty: {
                translate: 'transform',
            },
            transitionDuration: {
                'translate': '200ms',
            },
            keyframes: {
                "fade-in": {
                    "0%": { opacity: "0" },
                    "100%": { opacity: "1" }
                },
                "fade-out": {
                    "0%": { opacity: "1" },
                    "100%": { opacity: "0" },
                },
                "drop-down": {
                    "0%": { opacity: "0", transform: "translateY(5%)" },
                    "100%": { opacity: "1", transform: "translateY(0%)" },
                },
                "drop-up": {
                    "0%": { opacity: "1", transform: "translateY(0%)" },
                    "100%": { opacity: "0", transform: "translateY(5%)" },
                },
                "toast": {
                    "0%": {
                        opacity: 0,
                        transform: "translateX(-20%)",
                        scale: "0.95"
                    },
                    "50%": {
                        opacity: 1,
                        transform: "translateX(0)",
                        scale: "1"
                    },
                    "100%": {
                        opacity: 1,
                        transform: "translateX(0)",
                        scale: "1"
                    },
                },
                "drawer-open": {
                    "0%": {
                        opacity: 0.7,
                        transform: "translateX(30%)",
                    },
                    "100%": {
                        opacity: 1,
                        transform: "translateX(0)",
                    },
                },
                "drawer-close": {
                    "0%": {
                        opacity: 0.7,
                        transform: "translateX(0)",
                    },
                    "100%": {
                        opacity: 1,
                        transform: "translateX(30%)",
                    },
                },
                toastExit: {
                    '0%': {
                        opacity: '1',
                        transform: 'translateX(0) scale(1)',
                    },
                    '50%': {
                        opacity: '1',
                        transform: 'translateX(0) scale(1)',
                    },
                    '100%': {
                        opacity: '0',
                        transform: 'translateX(-20%) scale(0.95)',
                    },
                },
            },
            animation: {
                "fade-in": "fade-in 700ms ease-out",
                "fade-out": "fade-out 700ms ease-out",
                "toast": "toast 500ms ease-out",
                "toastExit": "toastExit 500ms ease-out",
                "drawer-open": "drawer-open 350ms ease-out",
                "drawer-close": "drawer-close 350ms ease-out",
                "drop-down": "drop-down 350ms ease-out",
                "drop-up": "drop-up 350ms ease-out",
            },
            gridTemplateColumns: {
                "edit-profile-form": "calc(100% / 3) 1fr",
                "create-service-packages-table": "10rem 1fr 1fr 1fr",
                service: "1fr 25rem",
                "sign-form": "36rem 1fr",
                "search-page": "1fr 23rem",
                "search-top-bar": "1fr auto",
                "orders-table": "5rem 6rem calc(100% - 38rem) 13rem 6rem 8rem",
                "services-table": "5rem 7rem calc(100% - 51.5rem) 13rem 4.5rem 8rem 6rem 8rem",
                "requests-table": "10rem 10rem calc(100% - 47rem) 13rem 6rem 8rem",
                "orders-table-mobile": "6rem 1fr",
                "order-overview-receipt": "8rem 1fr 5.5rem",
                "two-columns-with-divider": "1fr 1px 1fr",
            },
            spacing: {
                4.5: "1.125rem",
                packages: "calc(32px)",
                22: "5.5rem",
                10: "2.5rem",
                19: "4.75rem",
                15: "3.75rem",
            },
            width: {
                100: "25rem",
                112: "28rem",
                "search-sidebar": "23rem",
                48: "12rem",
                55: "13.75rem",
            },
            minWidth: {
                "header-dropdown-menu": "15rem",
                60: "15rem",
            },
            height: {
                form: "56rem",
                128: "32rem",
                "order-overview": "23.3125rem",
                10.5: "2.625rem",
                13.5: "3.375rem",
            },
            translate: {
                "search-sidebar": "23rem",
                112: "28rem",
            },
            maxWidth: {
                wrapper: "85rem",
                256: "64rem",
                168: "42rem",
            },
            minHeight: {
                app: "calc(100vh - 4rem - 3rem - 4rem - 1px)",
                212: "53rem",
                "search-left-side-mobile": "calc(100vh - 4rem - 3rem)",
            },
            maxHeight: {
                "search-sidebar": "calc(100vh - 4rem - 3rem)",
                app: "calc(100vh - 4rem - 3rem - 4rem - 1px)",
                "search-array": "calc(100vh - 4rem  - 3rem - 4rem - 6rem  - 3rem)",
                "search-array-mobile": "calc(100vh - 4rem - 6rem - 3rem)",
                "chat-body": "calc(100vh - 4rem - 3rem - 4rem - 1px - 3.25rem - 1.5rem - 5.5rem - 1.5rem - 6rem)",
                81.5: "20.375rem",
            },
            colors: {
                dark: {
                    palette: {
                        action: {
                            active: "#fff",
                            hover: "rgba(255, 255, 255, 0.08)",
                            selected: "rgba(255, 255, 255, 0.16)",
                            disabled: "rgba(255, 255, 255, 0.32)",
                            disabledBackground: "rgba(255, 255, 255, 0.12)",
                        },
                        text: {
                            primary: "#fff",
                            secondary: "rgba(255, 255, 255, 0.7)",
                            disabled: "rgba(255, 255, 255, 0.5)",
                        },
                        background: {
                            default: "#1a1a1a",
                            block: "#222222",
                        },
                        divider: "rgba(255, 255, 255, 0.2)",
                    }
                },
                light: {
                    palette: {
                        action: {
                            active: "#000",
                            hover: "rgba(0, 0, 0, 0.08)",
                            selected: "rgba(0, 0, 0, 0.16)",
                            disabled: "rgba(0, 0, 0, 0.32)",
                            disabledBackground: "rgba(0, 0, 0, 0.12)",
                        },
                        text: {
                            primary: "#000",
                            secondary: "rgba(0, 0, 0, 0.7)",
                            disabled: "rgba(0, 0, 0, 0.5)",
                        },
                        background: {
                            default: "#f9f9f9",
                            block: "#fff"
                        },
                        divider: "rgba(0, 0, 0, 0.2)",
                    }
                }
            }
        },
    },
    plugins: [],
    darkMode: "class",
};

export default config;
