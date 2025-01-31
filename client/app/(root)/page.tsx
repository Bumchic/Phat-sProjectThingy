"use client";
import React, { useEffect, useState } from "react";
import HeaderBox from "@/components/HeaderBox";
import TotalBalanceBox from "@/components/TotalBalanceBox";
import Navbar from "@/components/ui/navbar";
import Wallet from "@/components/Wallet";
import "../globals.css";
import { DetailService } from "@/components/DetailService";

export default function Home() {
  const [users, setUsers] = useState([]);
  const backendUrl = "http://localhost:8080";
  const loggedIn = { firstName: users[0] || "Guest" };

  useEffect(() => {
    fetch(`${backendUrl}/v1/user`)
      .then((response) => response.json())
      .then((data) => setUsers(data))
      .catch((error) => console.error("Error fetching users:", error));
  }, [backendUrl]);

  const handlePayNow = () => {};

  const handleHistory = () => {};

  return (
    <section className="home">
      <div className="home-content">
        <div className="home-header ">
          <Navbar />
          <div className="ml-8">
            <HeaderBox
              type="greeting"
              title="Welcome"
              user={loggedIn?.firstName || "Guest"}
              subtext="Access and manage your account"
            />
          </div>

          <div className="bg-white-50 flex justify-center p-8 ">
            <div className="flex justify-between w-full">
              <TotalBalanceBox
                accounts={[]}
                totalBanks={1}
                totalCurrentBalance={3000000}
              />
              <Wallet
                balance={70000}
                onPayNow={handlePayNow}
                onHistory={handleHistory}
              />
            </div>
          </div>

          <div className="w-100% p-8 bg-gray-50 flex flex-col rounded-lg">
            <h1 className="flex justify-center text-3xl font-bold text-gray-800 py-8 ">
              Your Service (Registered)
            </h1>
            <DetailService/>
          </div>
        </div>
      </div>
    </section>
  );
}
