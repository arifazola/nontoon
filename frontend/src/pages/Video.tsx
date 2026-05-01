import { MediaControlBar, MediaController, MediaFullscreenButton, MediaMuteButton, MediaPlaybackRateButton, MediaPlayButton, MediaSeekBackwardButton, MediaSeekForwardButton, MediaTimeDisplay, MediaTimeRange, MediaVolumeRange } from 'media-chrome/react'
import { MediaRenditionMenu, MediaRenditionMenuButton } from 'media-chrome/react/menu'
import { useEffect, useState } from 'react'
import ReactPlayer from 'react-player'
import { useParams } from 'react-router'

const Video = () => {
    const params = useParams()
    const [videoReady, setVideoReady] = useState<boolean | null>(null)
    useEffect(() => {
        const isVideoReady = async () => {
            const result = await fetch(`http://localhost:8080/videos/${params.id}`)
            const res = await result.json()

            setVideoReady(res.status)
        }

        isVideoReady()
    }, [])
    if (videoReady === null) {
        return (
            <p>Loading</p>
        )
    } else if (!videoReady) {
        return (
            <p>Your video is being processed</p>
        )
    } else {
        return (
            <MediaController
                style={{
                    width: "100%",
                    aspectRatio: "16/9",
                }}
            >
                <ReactPlayer
                    slot="media"
                    src={`http://localhost:8080/assets/${params.id}/master.m3u8`}
                    controls={false}
                    style={{
                        width: "100%",
                        height: "100%",
                    }}
                ></ReactPlayer>
                <MediaRenditionMenu anchor="auto" />
                <MediaControlBar>
                    <MediaPlayButton />
                    <MediaSeekBackwardButton seekOffset={10} />
                    <MediaSeekForwardButton seekOffset={10} />
                    <MediaTimeRange />
                    <MediaTimeDisplay showDuration />
                    <MediaRenditionMenuButton />
                    <MediaMuteButton />
                    <MediaVolumeRange />
                    <MediaPlaybackRateButton />
                    <MediaFullscreenButton />
                </MediaControlBar>
            </MediaController>
        )
    }

}

export default Video