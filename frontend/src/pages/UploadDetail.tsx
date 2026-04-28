import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router'
import type { VideoJobs } from '../models/VideoJobs'

const UploadDetail = () => {
    const [latestChunk, setLatestChunk] = useState(0)
    const [totalChunks, setTotalChunks] = useState(0)
    const [totalChunksUploaded, setTotalChunksUploaded] = useState(0)
    let params = useParams()

    useEffect(() => {
        const getLatestUploadedChunk = async () => {
            const latestUpload = await fetch(`http://localhost:8080/chunks/${params.id}`)

            const res = await latestUpload.json() as VideoJobs

            console.log("latest chunk", res)

            setLatestChunk(res.index)
        }

        getLatestUploadedChunk()
    }, [])

    const handleInputChange = async (e: React.ChangeEvent<HTMLInputElement, HTMLInputElement>) => {
        if (e.target.files && e.target.files.length > 0) {
            const file = e.target.files[0]
            const chunks = createChunks(file, 1024 * 1024)

            console.log("chunks length", chunks.length)

            for (let i = latestChunk + 1; i < chunks.length; i++) {
                const formData = new FormData()

                const f = blobToFile(chunks[i], file.name)

                formData.append("file", f)
                formData.append("uploadId", "def")
                formData.append("chunkIndex", i.toString())
                const response = await fetch("http://localhost:8080/videos/chunks", {
                    method: "POST",
                    body: formData
                })

                const res = await response.json()
                setTotalChunksUploaded(i)
                console.log("result api", res)
            }
        }
    }

    const blobToFile = (theBlob: Blob, fileName: string): File => {
        return new File([theBlob], fileName, {
            type: theBlob.type,
            lastModified: Date.now()
        });
    };

    const createChunks = (file: File, chunkSize: number): Blob[] => {
        const chunks: Blob[] = [];
        for (let start = 0; start < file.size; start += chunkSize) {
            const chunk = file.slice(start, start + chunkSize);
            chunks.push(chunk);
        }
        setTotalChunks(chunks.length)
        return chunks;
    };

    return (
        <>
            <div>UploadDetail</div>

            <br />

            {latestChunk > 0 ? (
                <p>Upload interupted at {latestChunk}</p>
            ) : null}

            <input type='file' onChange={(e) => handleInputChange(e)}></input>

            <br />

            <p>Total chunks uploaded: {totalChunksUploaded} out of {totalChunks}</p>
        </>


    )
}

export default UploadDetail