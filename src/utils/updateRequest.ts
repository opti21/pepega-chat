import { prisma } from "./prisma";

export async function updateRequest(
  requestID: number,
  videoID: number
): Promise<boolean> {
  try {
    await prisma.request.update({
      where: {
        id: requestID,
      },
      data: {
        video_id: videoID,
      },
    });

    return true;
  } catch (e) {
    console.error("Error updating request: ", e);
    return false;
  }
}
