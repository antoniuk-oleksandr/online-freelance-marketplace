import {createTheme, MantineThemeOverride} from '@mantine/core';

export const mantineTheme = (colorScheme: 'light' | 'dark'): MantineThemeOverride => {
    return createTheme({
        colors: {
            dark: [
                "#1a1a1a", // Background default
                "#222222", // Background block
                "rgba(255, 255, 255, 0.12)", // Divider
                "#fff", // Text primary
                "rgba(255, 255, 255, 0.7)", // Text secondary
                "rgba(255, 255, 255, 0.5)", // Text disabled
                "rgba(255, 255, 255, 0.08)", // Action hover
                "rgba(255, 255, 255, 0.16)", // Action selected
                "rgba(255, 255, 255, 0.32)", // Action disabled
                "rgba(255, 255, 255, 0.12)", // Action disabledBackground,
                "#ef4444" // Error
            ],
            light: [
                "#f9f9f9", // Background default
                "#fff", // Background block
                "rgba(0, 0, 0, 0.12)", // Divider
                "#000", // Text primary
                "rgba(0, 0, 0, 0.7)", // Text secondary
                "rgba(0, 0, 0, 0.5)", // Text disabled
                "rgba(0, 0, 0, 0.08)", // Action hover
                "rgba(0, 0, 0, 0.16)", // Action selected
                "rgba(0, 0, 0, 0.32)", // Action disabled
                "rgba(0, 0, 0, 0.12)", // Action disabledBackground
                "#ef4444" // Error,
            ],
        },
        components: {
            Input: {
                styles: (theme: any, {error}: { error: boolean}) => ({
                    input: {
                        height: "2.75rem",
                        fontSize: "1rem",
                        backgroundColor: colorScheme === 'dark' ? theme.colors.dark[0] : theme.colors.light[0],
                        color: error
                            ? (colorScheme === 'dark' ? theme.colors.dark[10] : theme.colors.light[10])
                            : (colorScheme === 'dark' ? theme.colors.dark[3] : theme.colors.light[3]),
                        borderColor: error
                            ? (colorScheme === 'dark' ? theme.colors.dark[10] : theme.colors.light[10])
                            : (colorScheme === 'dark' ? theme.colors.dark[2] : theme.colors.light[2]),
                        '&:hover': {
                            backgroundColor: colorScheme === 'dark' ? theme.colors.dark[6] : theme.colors.light[6],
                        },
                        '&:disabled': {
                            backgroundColor: colorScheme === 'dark' ? theme.colors.dark[9] : theme.colors.light[9],
                            color: colorScheme === 'dark' ? theme.colors.dark[5] : theme.colors.light[5],
                        },
                    }
                }),
            },
            PasswordInput: {
                styles: (theme: any, {error} : {error: boolean}) => ({
                    input: {
                        height: "2.875rem",
                        fontSize: "1rem",
                    },
                    visibilityToggle: {
                        color: error
                            ? (colorScheme === 'dark' ? theme.colors.dark[10] : theme.colors.light[10])
                            : (colorScheme === 'dark' ? theme.colors.dark[3] : theme.colors.light[3]),
                    },
                }),
            },
            InputWrapper: {
                styles: (theme: any, {error}: { error: boolean }) => ({
                    label: {
                        fontSize: "1rem",
                        color: error
                            ? (colorScheme === 'dark' ? theme.colors.dark[10] : theme.colors.light[10])
                            : (colorScheme === 'dark' ? theme.colors.dark[3] : theme.colors.light[3]),
                    },
                }),
            },
        },
    })
}