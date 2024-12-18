<script lang="ts">
    import DropdownMenu from "@/common-components/Dropdown/components/DropdownMenu/DropdownMenu.svelte";
    import type {DropdownProps} from "@/types/DropdownProps.ts";
    import {resolveDropdownAnimation} from "@/common-components/Dropdown/helper.ts";
    import DropdownLayout from "@/common-components/Dropdown/DropdownLayout.svelte";
    import DropdownTrigger from "@/common-components/Dropdown/components/DropdownTrigger/DropdownTrigger.svelte";

    const props: DropdownProps = $props();

    let isOpen: boolean = $state(false);
    const setIsOpen = (value: boolean) => isOpen = value;

    let showExitAnimation: boolean = $state(false);
    const setShowExitAnimation = (value: boolean) => showExitAnimation = value;

    let timeout: number | undefined;
    const setTimeoutValue = (value: number | undefined) => timeout = value;

    $effect(() => resolveDropdownAnimation(
        timeout, setTimeoutValue,
        isOpen, setShowExitAnimation
    ));
</script>

<DropdownLayout>
    <DropdownTrigger
            dropdownMenuProps={props}
            title={props.title}
            isOpen={isOpen} setIsOpen={setIsOpen}
    >
        {@render props.children()}
    </DropdownTrigger>
    <DropdownMenu
            showExitAnimation={showExitAnimation}
            isOpen={isOpen}
            setIsOpen={setIsOpen}
            {...props}
    />
</DropdownLayout>
