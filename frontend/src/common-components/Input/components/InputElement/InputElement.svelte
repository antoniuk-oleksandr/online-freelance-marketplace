<script lang="ts">
    type InputElementProps = {
        setIsFocused: (value: boolean) => void,
        id: string,
        type: "text" | "password",
    }

    const {id, type, setIsFocused}: InputElementProps = $props();
</script>

<input
        onfocus={() => setIsFocused(true)}
        onblur={() => setIsFocused(false)}
        class="p-4 w-full outline-none bg-transparent text-lg text-light-palette-text-primary dark:text-dark-palette-text-primary"
        id={id} name={id} type={type}
>