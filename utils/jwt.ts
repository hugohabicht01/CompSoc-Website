import type { PrismaClient } from "@prisma/client";
import jwt from "jsonwebtoken";
import { jwtVerify as joseJwtVerify } from "jose";

const { secret } = useRuntimeConfig();
const encodedSecret = new TextEncoder().encode(secret);

export type jwtDecodedType = {
  iat: number;
  exp: number;
  username: string;
  displayName: string;
  mail: string;
};

export type Role = "USER" | "PRIVILEGED" | "ADMIN";
export type jwtPayloadType = {
  username: string;
  displayName: string;
  mail: string;
  banned: boolean;
  role: Role;
};

export async function verifyIssJwt(jwt: string): Promise<jwtDecodedType> {
  const { payload } = await joseJwtVerify(jwt, encodedSecret);
  return payload as unknown as jwtDecodedType;
}

export async function verifyJWT(token: string) {
  return jwt.verify(token, secret) as jwtPayloadType & {
    iat: number;
    exp: number;
  };
}

export async function createJWT(
  issJwtData: jwtDecodedType,
  prisma: PrismaClient,
) {
  const { username } = issJwtData;
  const user = await prisma.user.findUnique({ where: { username } });
  if (!user) {
    throw new Error("user not found");
  }

  const payload: jwtPayloadType = {
    username,
    banned: user.banned,
    role: user.role,
    displayName: user.displayName,
    mail: user.mail,
  };

  return jwt.sign(payload, secret, {
    expiresIn: "1 week",
    issuer: "compsoc",
  });
}
