//
//  Config.swift
//  MomentumIOS
//
//  Created by Alder Whiteford on 8/13/24.
//

import Foundation
import Supabase

let supabaseKey: String = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImNrb3hubnlvYXVjaHV4cG5pcHNjIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjM0ODkzOTUsImV4cCI6MjAzOTA2NTM5NX0.b1z1rOK6vqqL5shpz49BOmZdXTO9PXdW87lmdfHufN4"

let supabaseURL: URL = URL(string: "https://ckoxnnyoauchuxpnipsc.supabase.co")!

let supabase = SupabaseClient(
    supabaseURL: supabaseURL,
    supabaseKey: supabaseKey
)
