import React from 'react'
import Link from "next/link";
// import Login from '../(auth)/login/page';

import "../globals.css";
export default function Home() {
  return (
    <div>
      <h2 className="text-red-500 ">Welcome to Regular Billing</h2>
      <Link href="/login" className="text-red-500">
        Go to Login
      </Link>

      <br />
      <Link href="/signup" className="text-red-500">
        Go to Sign Up
      </Link>
      <br />
      <Link href="/mainpage" className="text-red-500">
        Go to Main page
      </Link>
    </div>
  );
}
