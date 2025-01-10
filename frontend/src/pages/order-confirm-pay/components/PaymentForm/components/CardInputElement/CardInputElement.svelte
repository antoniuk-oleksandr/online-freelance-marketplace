<script lang="ts">
    type CardInputElementProps = {
        id: string,
        placeholder: string,
        error?: string,
        focusAction?: () => void,
        blurAction?: () => void,
    }

    let props: CardInputElementProps = $props();
</script>

<input
        autocomplete="off"
        class="{props.error ? '!ring-red-500' : ''} bg-light-palette-background-block dark:bg-dark-palette-background-block ring-1 ring-light-palette-divider dark:ring-dark-palette-divider focus:ring-cyan-500 focus:dark:ring-cyan-500 duration-200 ease-out rounded-md py-2 px-3 outline-none w-full"
        id={props.id}
        name={props.id}
        onblur={props.blurAction}
        onfocus={props.focusAction}
        placeholder={props.placeholder}
        type="text"
/>