<template>
    <button type="button" class="button">
        <span class="button__text">{{ text }}</span>
        <span class="button__icon"><svg xmlns="http://www.w3.org/2000/svg" width="24" viewBox="0 0 24 24" stroke-width="2"
                stroke-linejoin="round" stroke-linecap="round" stroke="currentColor" height="24" fill="none" class="svg">
                <line y2="19" y1="5" x2="12" x1="12"></line>
                <line y2="12" y1="12" x2="19" x1="5"></line>
            </svg></span>
    </button>
</template>

<script lang="ts" setup>
defineProps({
    text: {
        type: String,
        default: "Add Item"
    }
})
</script>

<style lang="less" scoped>
@var1: #34974d;
@var2: #3aa856;
@var3: #ffffff;
@var4: #2e8644;

.button {
    position: relative;
    width: 150px;
    height: 40px;
    cursor: pointer;
    display: flex;
    align-items: center;
    border: 1px solid @var1;
    background-color: @var2;

    .button__text {
        // transform: translateX(30px);
        color: @var3;
        font-weight: 600;
        width: calc(100% - 39px);
    }

    .button__icon {
        position: absolute;
        transform: translateX(109px);
        height: 100%;
        width: 39px;
        background-color: @var1;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .svg {
        width: 30px;
        stroke: @var3;
    }

    &:hover {
        background: @var1;

        .button__text {
            color: transparent;
        }

        .button__icon {
            width: 148px;
            transform: translateX(0);
        }
    }

    &:active .button__icon {
        background-color: @var4;
    }

    &:active {
        border: 1px solid @var4;
    }
}

.button,
.button__icon,
.button__text {
    transition: all 0.3s;
}
</style>