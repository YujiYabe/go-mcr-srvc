<template>
	<div>
		<div v-if="!isOnline" align="center" class="wrapper">
			<icon-reloading> </icon-reloading>
		</div>
		<div v-else>
			<b-container fluid class="text-light text-center">
				<b-row>
					<b-col
						cols="6"
						:class="[commonClass, methodClassSelected('bed')]"
						@click="methodChangeRoom('bed')"
					>
						<icon-room-bed />
					</b-col>

					<b-col
						cols="6"
						:class="[commonClass, methodClassSelected('living')]"
						@click="methodChangeRoom('living')"
					>
						<icon-room-living />
					</b-col>
				</b-row>
			</b-container>

			<!-- ボディ -->
			<div v-show="room == 'bed'">
				<bed
					:device="device"
					:common-class="commonClass"
					:light-status="lightStatus"
					:aircon-status="airconStatus"
					:aircon-warm="airconWarm"
					:aircon-cool="airconCool"
					@change-device="methodChangeDevice"
					@send-command="methodSendCommand"
				/>
			</div>

			<div v-show="room == 'living'">
				<living
					:device="device"
					:common-class="commonClass"
					:light-status="lightStatus"
					:aircon-status="airconStatus"
					:aircon-warm="airconWarm"
					:aircon-cool="airconCool"
					:files="files"
					:vlc-progress="vlcProgress"
					:is-downloading="isDownloading"
					:is-uploading="isUploading"
					@change-device="methodChangeDevice"
					@send-command="methodSendCommand"
					@file-upload="methodFileUpload"
				/>
			</div>
		</div>
	</div>
</template>

<script>
module.exports = {
	props: {
		'commonClass': { type: String, default: '' },
		'room': { type: String, default: '' },
		'device': { type: String, default: '' },
		'deviceBed': { type: String, default: '' },
		'deviceLiving': { type: String, default: '' },

		'lightStatus': { type: String, default: '' },
		'airconStatus': { type: String, default: '' },
		'airconWarm': { type: Number, default: 0 },
		'airconCool': { type: Number, default: 0 },

		'files': { type: Array, default: [] },
		'vlcProgress': { type: Number, default: 0 },
		'isDownloading': { type: Boolean, default: false },
		'isUploading': { type: Boolean, default: false },
		'isOnline': { type: Boolean, default: false },
	},

	methods: {
		methodChangeRoom: function (value) {
			this.$emit('change-room', value);
		},

		methodChangeDevice(value) {
			this.$emit('change-device', value);
		},

		methodSendCommand: function (value) {
			this.$emit('send-command', value);
		},

		methodFileUpload: function (value) {
			this.$emit('file-upload', value);
		},

		methodClassSelected(value) {
			return value == this.room ? 'btn-primary' : 'btn-secondary';
		},
	},

	components: {
		'icon-room-living': httpVueLoader('public/vue/icon/room-living.vue'),
		'icon-room-bed': httpVueLoader('public/vue/icon/room-bed.vue'),
		'icon-reloading': httpVueLoader('public/vue/icon/common-reloading.vue'),

		bed: httpVueLoader('public/vue/room/bed.vue'),
		living: httpVueLoader('public/vue/room/living.vue'),
	},
};
</script>
