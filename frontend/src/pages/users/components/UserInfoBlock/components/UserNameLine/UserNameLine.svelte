<script lang="ts">
    import type {UserByIdProps} from "@/types/UserByIdProps.ts";

    type UserNameLineProps = UserByIdProps & {
        size?: "small" | "large",
        useLink?: boolean,
    }

    const {user, size, useLink}: UserNameLineProps = $props();
</script>

<h2
        class={`font-semibold
        ${size === "small" ? "text-lg" : "text-3xl"}
        ${useLink ? "cursor-pointer hover:underline duration-200 ease-out" : ""}`}
>{user.firstName} {user.surname} @{user.username}</h2>